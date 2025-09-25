# Il nostro programma in esecuzione mostra i 5 job che vengono eseguiti da
# vari worker. Il programma impiega solo circa 2 secondi
# nonostante faccia circa 5 secondi di lavoro totale perché
# ci sono 3 worker che operano simultaneamente.
$ time go run pool-di-worker.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5

real	0m2.358s
