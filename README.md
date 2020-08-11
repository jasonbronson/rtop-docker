# Modified Rtop which on route shows you stats of your docker containers and host machine resource uses

Runs rtop project in a docker container that on route / will return json data output of machine resources

Example 

{
    "Uptime": "2d 13h 24m 23s",
    "Hostname": "mymachine",
    "Load1": "0.59",
    "Load5": "0.40",
    "Load10": "0.35",
    "RunningProcs": "1",
    "TotalProcs": "551",
    "MemTotal": "2.86 GiB",
    "MemUsed": "1.19 GiB",
    "MemFree": "448.59 MiB",
    "MemBuffers": "436.40 MiB",
    "MemCached": "829.49 MiB",
    "SwapTotal": "2.93 GiB",
    "SwapFree": "2.91 GiB",
    "FSInfos": null,
    "NetIntf": null,
    "CPU": {
        "User": 0.0862069,
        "Nice": 0,
        "System": 0.0862069,
        "Idle": 99.82759,
        "Iowait": 0,
        "Irq": 0,
        "SoftIrq": 0,
        "Steal": 0,
        "Guest": 0
    }
}