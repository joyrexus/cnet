This package provides a simple function and command-line utility for authenticating UChicago [CNetID](http://answers.uchicago.edu/page.php?id=16157) users.

To install run `go get`:

    $ go get github.com/joyrexus/cnet/...

This will retrieve the library and install a `cnet` command into your `$GOBIN` path.

To test your installation, try authenticating with your CNet user ID and password:

    $ cnet --user rzimmer --password ******

The package has a few dependencies, but will be installed automagically:
* [`gopkg.in/ldap.v1`](http://gopkg.in/ldap.v1)
* [`gopkg.in/asn1-ber.v1`](http://gopkg.in/asn1-ber.v1)

For info on UChicago's LDAP authentication services, see ...
* [overview](http://answers.uchicago.edu/page.php?id=16242)
* [authentication](http://answers.uchicago.edu/page.php?id=16242)

I'd like to use this as one piece to RESTful API service requiring
authentication on protected resources.  For more on how that could be done, see
[this overview](http://restcookbook.com/Basics/loggingin/) and [this
article](https://stormpath.com/blog/secure-your-rest-api-right-way/).
