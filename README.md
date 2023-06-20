# Macro Keyboard

Just another custom macro keyboard fully written in Go ([TinyGo](https://tinygo.org/) to be more specific).

## Hardware

An example of build was made using a [Seeeduino Xiao](https://wiki.seeedstudio.com/Seeeduino-XIAO/) by Seeed Studio and 10 switches.

All switches were wired directly to each input (with a pull-up) since it was enough for my test case, using key matrices are not supported at the moment.

## Flashing

To flash this to a controller you will need:
- [Go 1.19+](https://go.dev/doc/install)
- [Tinygo 0.28.1](https://tinygo.org/getting-started/install/)

Then at the root of the project run:
```
tinygo flash -target=xiao -monitor
```

## Configuration

Most the configuration are made on the files:
- `configs\configs.go`:
    - `PollingDelay`: define time in between polling;
    - `AllowRepeat`: define if the input will be repeated when held;
    - `RepeatDelay`: define time in between repeats;
    - `Buttons`: default settings for buttons (will override persistent storage on failures), look at [Action chains](#action-chains);
    - `ResetPin`: pin to reset persistent storage to default state (from `Buttons`).
- `internal\utils\pinmap.go`: from **pin name** to `machine.Pin` equivalent, must be changed only if using alternate controller or buttons pin map.

Once the firmware has been flashed, the serial port keeps monitoring for commands to add new shortcuts/action chains, below is an example on how to setup.

```
tinygo monitor -target=xiao  # connect to serial port
add D10::text##Hello world
```

This will configure the D10 button to type "Hello world" when pressed.

## Action chains

Below are a few examples of action settings:

- Keycode actions: will press the keycodes together before releasing as described by the format `keycode##<first_keycode>##<second_keycode>`
    - `keycode##KeyModifierCtrl##KeyC`: copy shortcut
    - `keycode##KeyModifierCtrl##KeyX`: cut shortcut
    - `keycode##KeyModifierCtrl##KeyV`: paste shortcut
- Mouse actions: move mouse and click with specified button as described by the format `mouse##<x>##<y>##<button>`
    - `mouse##0##0##left`: left click without moving the mouse;
    - `mouse##5##0##`: move mouse 5 pixels to the right;
    - `mouse##0##-10##right`: move mouse 10 pixels to the bottom and click right button;
- Text actions: type a text as described by the format `text##<text to type>`
- Delay actions: wait before doing next action in chain as described by the format `delay##<time with unit>`
    - `delay##50ms`: will wait for 50ms
    - `delay##1s`: will wait for 1s

Multiple actions can be described using the separator `$$`, they will be run in order, the example below will move the mouse to the right, wait 50ms and them move the mouse back to the left.
```
mouse##5##0##$$delay##50ms$$mouse##-5##0##
```
