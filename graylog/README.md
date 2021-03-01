What is this?
----

This is a set of scripts containing docker containers with graylog mongo and
elasticsearch to provide you with a easier way to read logs from log00 machines.

Every time there is a problem, be it in the TCs or in any environment that
doesn't have Splunk, we poor developers, are asked to see the log files and to
use freaking `grep` to search for information. If this was a decade (or two)
ago I guess it would be fine... If...

However we are in the 21st century and there are more efficient means of
getting information from logs. Means that don't make us blind as a side effect,
for example.

Graylog is an application that helps just with that. Logs. It is a competitor
of Splunk and has a community version, which is the one this set of scripts
has.

Installation
----

First you need to have docker installed. Then you simply run:

`./setup_garylog.sh`

This command will download, run and do any machine specific configurations for
you to run Graylog.

Initial Setup
----

After the installation you need to do login with:

```
   user: admin
   password: admin
```

And then do some configurations. The first thing to do is to add an Input:

1. System -> Inputs
2. Select "Raw/Plaintext TCP" from the select input dropbox
3. Click "Launch new input"
4. Click the "Global" checkbox
5. Give it the title "Telnet"
6. Scroll to the bottom and click "Save"

Now that you have an input, let's import some extractors:

1. Click "Manage extractors" button
2. Go to "Actions" -> "Import extractors"
3. Copy the contents of the `extractors.json` into the text area
4. Click "Add extractors to input"

You're done!
Happy logging !

How to use
----

Once you are done with the initial setup, it is recommended you power down
Graylog, mainly because it uses some ports that conflict with other projects.

To stop Graylog you can run the `./stop_graylog.sh` script.
To start it, you can run the `./start_graylog.sh` script.

Sometimes when starting Graylog, the graylog container starts before the
elasticsearch one is active. When this happens graylog crashes. To fix it
simply go to you Docker dashboard and ask the graylog container to run again.

Watching logs
---

This bundle has two main ways to get logs:

1. From the log00 machine: `./watch.sh`
2. From the jounalctl of a specific machine: `./follow_journal.sh`

To learn how to use them I recommend you check the script files and read the
documentation each one has.

Additionally, I also offer a `delete_all_indices.sh` which I recommend you use
after you are done with graylog and before stopping it, so you clear the logs
and avoid seeing repeated information next time you run `./watch.sh`.
