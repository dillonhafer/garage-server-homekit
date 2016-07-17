Garage Server HomeKit
------

A HomeKit server for Raspberry Pi to open a garage door. An alternative to [garage-server](https://github.com/dillonhafer/garage-server).

Hardware I used for project:

1. [Magnetic Reed Switch](http://amzn.to/1XuUrV9) (Optional. Used for door status)
2. [Relay Shield Module](http://amzn.to/1NRZf1R)

I really like the above relay because when the power is disconnected and restored *(i.e. power goes out in the middle of the night)* the relay will remain off. That way a power outage won't open your garage door.

## Options

```
  -pin string
    8-digit Pin for securing garage door
  -relay-pin int
    GPIO pin of relay (default 25)
  -sleep int
    Time in milliseconds to keep switch closed (default 500)
  -status-pin int
    GPIO pin of reed switch (default 10)
  -version
    print version and exit
```

## Installation Instructions

#### Installation Steps Overview:

1. **[Download garage-server-homekit](#user-content-download-garage-server-homekit)**
2. **[Create init.d script](#user-content-create-initd-script)**

#### Download garage-server-homekit

**Install from source**

Make sure [go](https://golang.org/) is installed on your Raspberry Pi and then you can use `go get` for installation:

```bash
go get github.com/dillonhafer/garage-server-homekit
```

**Install from binary**

Latest binaries available at https://github.com/dillonhafer/garage-server-homekit/releases/latest

#### Create init.d script

Simply copy the init.d script from the src directory.

```bash
cp $GOPATH/src/github.com/dillonhafer/garage-server-homekit/garage-server-homekit.init /etc/init.d/garage-server-homekit
```

#### Configure init.d script

The last thing to do is to configure your init.d script to reflect your Raspberry Pi's configuration.

First set the `GARAGE_SECRET` environment variable. This will ensure JSON requests to the server are authenticated. Be sure to use a very random and lengthy secret.

Just un-comment the following line and add your pin in the init.d script:

```bash
# /etc/init.d/garage-server-homekit...

# Remember to set a different 8-digit pin(e.g. 12391123)
# DO NOT USE the above pin. It's an example only.
PIN=12391123
```

Other configuration variables to consider are the `STATUS_PIN` and `RELAY_PIN`. Use these
to set what GPIO pins your Raspberry Pi is configured to use.

```bash
# /etc/init.d/garage-server-homekit...
STATUS_PIN=10
RELAY_PIN=25
```

Now just install and start the service:

```bash
sudo chmod +x /etc/init.d/garage-server-homekit
sudo update-rc.d garage-server-homekit defaults
sudo service garage-server-homekit start
```

That's it! The server is now setup!

## Notes

I would like to find a license-free implementation of the HomeKit bridge to allow this
project to be freely distributed in the future. I apologize for the current
inconvenience.

## Un-License

  This is free and unencumbered software released into the public domain.

  Anyone is free to copy, modify, publish, use, compile, sell, or
  distribute this software, either in source code form or as a compiled
  binary, for any purpose, commercial or non-commercial, and by any
  means.

  In jurisdictions that recognize copyright laws, the author or authors
  of this software dedicate any and all copyright interest in the
  software to the public domain. We make this dedication for the benefit
  of the public at large and to the detriment of our heirs and
  successors. We intend this dedication to be an overt act of
  relinquishment in perpetuity of all present and future rights to this
  software under copyright law.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
  IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
  OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
  ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
  OTHER DEALINGS IN THE SOFTWARE.

  For more information, please refer to <http://unlicense.org>

### This software uses `github.com/brutella/log` - see brutella-log-license.txt

### This software uses `github.com/brutella/hc` - see brutella-hc-license.txt
