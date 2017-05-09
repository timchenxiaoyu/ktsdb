package main

import (
	"time"
	"ktsdb/model"
	"ktsdb/chunk"
	"ktsdb/storage"
	"fmt"
)

func main() {
	chunk.DefaultEncoding = 2

	o := &storage.MemorySeriesStorageOptions{
		MemoryChunks:               1000000,
		MaxChunksToPersist:         1000000,
		PersistenceRetentionPeriod: 24 * time.Hour * 365 * 100, // Enough to never trigger purging.
		PersistenceStoragePath:    "/tmp/prom",
		CheckpointInterval:         time.Hour,
		SyncStrategy:               storage.Adaptive,
	}

	s := storage.NewMemorySeriesStorage(o)
	samples := make([]*model.Sample, 100)
	fingerprints := make(model.Fingerprints, 100)


	if err := s.Start(); err != nil {
		println("Error creating storage: %s", err)
	}

	for i := range samples {
		metric := model.Metric{
			model.MetricNameLabel: model.LabelValue(fmt.Sprintf("test_metric_%d", i)),
			"label1":              model.LabelValue(fmt.Sprintf("test_%d", i/10)),
			"label2":              model.LabelValue(fmt.Sprintf("test_%d", (i+5)/10)),
			"all":                 "const",
		}
		samples[i] = &model.Sample{
			Metric:    metric,
			Timestamp: model.Time(i),
			Value:     model.SampleValue(i),
		}
		fingerprints[i] = metric.FastFingerprint()
	}
	for i, a := range samples {
		println(i)
		s.Append(a)
	}
	s.WaitForIndexing()

}
