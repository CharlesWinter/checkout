# Checkout ![Test](https://github.com/CharlesWinter/checkout/workflows/Test/badge.svg?branch=main)

The checkout package provides functionality for scanning items, and telling the
user the price after any deals are applied.

*Note!* The task description didn't list any errors in the interface this
package had to adhere to, so I don't return any. Needless to say in a real
application, I would be returning errors.

## Test it
`make gen_mocks && make test`

## View the Documentation
Because the repository is public, the documentation should be available at.

https://godoc.org/github.com/CharlesWinter/checkout/checkout

The example of how the package should be used in practise can be seen at the
top of the page above, before "Index" by clicking the "Example" dropdown.

## Code Structure

I've tried to provide an example of how I would approach a problem like this in
the real world. There are some obvious deficiencies with the methodology as it
stands (e.g. deals are applied in the order the deal repo returns them, which needs
to be accounted for).

Rather than focusing on working out all the kinks in the system, I've chosen to
focus on demonstrating the following points, which I always like to see in code
and are largely taken from a book I reference a lot in my work; "Clean
Architecture" by Rob C Martin.

### Proper Abstractions:
Two obvious abstractions present themselves immediately; the pricing and the
deals. The "business logic" that provides value (in our case the calculation of
price for a basket) doesn't need to care where the deals and pricing comes
from. It could be from this very same program (which it is at the moment!), a
DB or another service. Through providing this abstraction, we gain two obvious
advantages straight away.

* We are free to change the _source_ of the pricing and deals at any point, and
  the checkout logic and tests remain exactly the same!

* Our testing becomes much easier, as these dependencies can be tested in their
  own packages and mocked in the checkout package, allowing us to focus on
  testing proper "units" i.e. business domains independently.

It could be said that this is premature optimisation, however I've always found
it much, much easier to move from abstractions to concretions than vice-versa
so I don't mind using some interfaces from the get go!

### Dependency Inversion
Following on from the above, interfaces are defined with the _consumer_ not the
_provider_. This further decouples the "business logic" from the lower level
concerns. We can think of the Checkout Repository as saying "I don't care who
you are, so long as you have these particular methods, I'll call them". It also
means the two packages are completely decoupled (notice they aren't in each
others `imports`) and are tied together through the domain entities and
eventually at the main when the `deals` and `prices` will be injected into the
checkout.

### Testing through public interfaces
A small point, but declaring your test packages with the format `package
checkout_test` in our case only makes the public interface visible.

## Things to Improve
If this were a real application, the following remains to be done.

- Add more tests (only happy path tests at the moment)
- Implement some sort of "deal priority" for where multiple deals are
  applicable.
- Implement error handling (this only isn't here because of the task
  description not specifying return errors!)
