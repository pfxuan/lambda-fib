This repo is used to support [KLR](https://github.com/triggermesh/knative-lambda-runtime) performance study across multiple languages through a series of load tests.


### go-1.x

1. Deploy function

```
cd go-1.x
tm deploy service go-lambda-fib -f . --build-template knative-go-runtime --wait
```

2. Run a simple test

```
curl <SERVICE_ENDPOINT> --data '{"Number": 20}'
{"id:":"d2809930-cec8-484a-a330-c80946a9fbf3","computeTime:":59010}
```

3. Perform load tests
```
npm install -g loadtest
loadtest -m POST -n 765 -c 1 --data '{"number":1}' <SERVICE_ENDPOINT>
INFO Requests: 0 (0%), requests per second: 0, mean latency: 0 ms
INFO Requests: 558 (73%), requests per second: 112, mean latency: 8.9 ms
INFO
INFO Target URL:          <SERVICE_ENDPOINT>
INFO Max requests:        765
INFO Concurrency level:   1
INFO Agent:               none
INFO
INFO Completed requests:  765
INFO Total errors:        0
INFO Total time:          6.896580553 s
INFO Requests per second: 111
INFO Mean latency:        8.9 ms
INFO
INFO Percentage of the requests served within a certain time
INFO   50%      7 ms
INFO   90%      11 ms
INFO   95%      13 ms
INFO   99%      25 ms
INFO  100%      122 ms (longest request)

loadtest -m POST -n 19125 -c 25 --data '{"number":1}' <SERVICE_ENDPOINT>
INFO Requests: 0 (0%), requests per second: 0, mean latency: 0 ms
INFO Requests: 4171 (22%), requests per second: 836, mean latency: 29.8 ms
INFO Requests: 7830 (41%), requests per second: 732, mean latency: 33.9 ms
INFO Requests: 9859 (52%), requests per second: 406, mean latency: 61.8 ms
INFO Requests: 13882 (73%), requests per second: 805, mean latency: 31.1 ms
INFO Requests: 17918 (94%), requests per second: 807, mean latency: 30.9 ms
INFO
INFO Target URL:          <SERVICE_ENDPOINT>
INFO Max requests:        19125
INFO Concurrency level:   25
INFO Agent:               none
INFO
INFO Completed requests:  19125
INFO Total errors:        0
INFO Total time:          26.492010719000003 s
INFO Requests per second: 722
INFO Mean latency:        34.5 ms
INFO
INFO Percentage of the requests served within a certain time
INFO   50%      24 ms
INFO   90%      56 ms
INFO   95%      84 ms
INFO   99%      206 ms
INFO  100%      931 ms (longest request)

loadtest -m POST -n 38250 -c 50 --data '{"number":1}' <SERVICE_ENDPOINT>
INFO Requests: 0 (0%), requests per second: 0, mean latency: 0 ms
INFO Requests: 4058 (11%), requests per second: 813, mean latency: 61.5 ms
INFO Requests: 8192 (21%), requests per second: 827, mean latency: 60.3 ms
INFO Requests: 12408 (32%), requests per second: 843, mean latency: 59.4 ms
INFO Requests: 16654 (44%), requests per second: 849, mean latency: 58.8 ms
INFO Requests: 20782 (54%), requests per second: 825, mean latency: 60.6 ms
INFO Requests: 25020 (65%), requests per second: 848, mean latency: 57.3 ms
INFO Requests: 29084 (76%), requests per second: 813, mean latency: 62.3 ms
INFO Requests: 33359 (87%), requests per second: 855, mean latency: 59.1 ms
INFO Requests: 37418 (98%), requests per second: 812, mean latency: 61.7 ms
INFO
INFO Target URL:          <SERVICE_ENDPOINT>
INFO Max requests:        38250
INFO Concurrency level:   50
INFO Agent:               none
INFO
INFO Completed requests:  38250
INFO Total errors:        0
INFO Total time:          46.042780053 s
INFO Requests per second: 831
INFO Mean latency:        60.1 ms
INFO
INFO Percentage of the requests served within a certain time
INFO   50%      48 ms
INFO   90%      103 ms
INFO   95%      142 ms
INFO   99%      263 ms
INFO  100%      1123 ms (longest request)
```

### node-10.x
[TODO]

### node-4.x
[TODO]

### python-2.7
[TODO]
### python-3.7
[TODO]
### ruby-2.5
[TODO]