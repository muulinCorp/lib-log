# log

## fluentd's spec in .env or env.yml
```yaml
fluentd:
    fluentd:
      host: fluentd
      port: 24224
```

## other's var
### .env
```env
LOG_LEVEL: debug/info/warn/error
LOG_TARGET: os/fluentd/os+fluentd
```
### env.yml
```yaml
log_level: debug/info/warn/error
log_target: os/fluentd/os+fluentd
```

## note
if var is not given:
  - log's level default value will be info
  - log's target default value will be os