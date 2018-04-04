# kotori-ng-sampleplugin

This plugin demonstrates how to develop a plugin
for the blog framework [kotori-ng](https://github.com/cool2645/kotori-ng).

To give a brief introduction of the framework itself, it is a web server
designed highly-plugable, which might be running attached with a bundle of plugins.

This repository will give you a sense of how to develop a plugin.

Here are what a plugin can do:

+ Define its own configuration files.
+ Define its own routes, under a specified subroute.
+ Access to the database instance of the kotori framework
(using the package [github.com/cool2645/kotori-ng/database](https://github.com/cool2645/kotori-ng/tree/master/database))
+ Use the JWT based auth session provided by the framework
(using the package [github.com/cool2645/kotori-ng/auth](https://github.com/cool2645/kotori-ng/tree/master/auth))
+ Use some convenient functions provided by the framework, like
[github.com/cool2645/kotori-ng/httputils](https://github.com/cool2645/kotori-ng/tree/master/httputils)

Furthermore, look though the code, it is very simple and easy to start with.