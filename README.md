# Brightness control

Change brightness value on Linux devices

## CLI interface

```
usage of brightnessctl:
  -abs
    	set value as is instead of inc/dec by
  -get
    	get current value
  -percent
    	use percents instead of absolute value
  -value int
    	brightness value (default 10)
```

### Usage

- `brightnessctl -value 10 --percent` - Increment value by 10% 
- `brightnessctl -value -17 --percent` - Decrement value by 17% 
- `brightnessctl -value 400 -abs` - Set specific value
- `brightnessctl -value 50 -percent -abs` - Set brightness half of max value

