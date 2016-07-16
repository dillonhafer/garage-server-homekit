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

1. **[Download garage-server-homekit](#download-garage-server-homekit)**
2. **[Create init.d script](#create-initd-script)**

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

## License

   Copyright 2016 Dillon Hafer

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

## Addendum License

   The vendor library `github.com/brutella/hc` has the below license:

   Copyright (c) 2015, Matthias Hochgatterer
   All rights reserved.

   Redistribution and use in source and binary forms,
   with or without modification, are permitted provided
   that the following conditions are met:

   - Redistributions of source code must retain the above copyright notice,
     this list of conditions and the following disclaimer.

   - Any redistribution, use, or modification is done solely for personal
     benefit and not for any commercial purpose or for monetary gain.

   THIS SOFTWARE IS PROVIDED BY THE PROJECT AND CONTRIBUTORS ''AS IS'' AND
   ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
   IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
   ARE DISCLAIMED.  IN NO EVENT SHALL THE PROJECT OR CONTRIBUTORS BE LIABLE
   FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
   DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
   OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
   HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
   LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
   OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
   SUCH DAMAGE.
