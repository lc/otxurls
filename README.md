# otxurls
Fetch known URLs from AlienVault's [Open Threat Exchange](https://otx.alienvault.com) for given hosts.

### usage:
```
▻ printf 'example.com' | otxurls
```

or

```
▻ otxurls example.com
```

### install:
```
▻ go get github.com/lc/otxurls
```


### Docker

Build
```
docker build -t otxurls .
```

Run
```
docker run --rm -t otxurls <url>
```
