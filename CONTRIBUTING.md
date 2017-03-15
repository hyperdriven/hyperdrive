# Contributing

First of all, thank you for wanting to contibute to a Hyperdriven project! :raised_hands:

We welcome and encourage contibutions (no matter how big or small) from **everyone**. By participating, you agree to abide by [our code of conduct](https://hyperdriven.net/code-of-conduct/).

## Ways You Can Contribute

  - [Submit a Pull Request](#pull-requests) :octocat: :shipit:
  - [Edit the Wiki](https://github.com/hyperdriven/hyperdrive/wiki) :pencil:
  - [Report a Bug](https://github.com/hyperdriven/hyperdrive/issues/new?labels=bug) :beetle:
  - [Suggest a Feature](https://github.com/hyperdriven/hyperdrive/issues/new?labels=enhancement) :bulb:
  - [Get Involved in The Community](https://hyperdriven.net/community/) :busts_in_silhouette:
  - [Tell us](mailto:hello@hypedriven.net) about your _hyperdriven API_. :mailbox_with_mail:
  - Write about us on your blog or social media! :mega:

## Pull Requests

1. [Fork the repo](https://github.com/hyperdriven/hyperdrive#fork-destination-box)

2. Clone your fork:

    git clone git@github.com:your-username/hyperdrive.git

3. Set up your machine:

    ./scripts/setup

4. Install dependencies:

    ./scripts/install

5. Make sure the tests pass:

    ./scripts/test

6. Make your change. Add tests, and ensure the tests pass:

    ./scripts/test

7. Commit your changes (squash your commits, if needed), push your fork, and [submit a pull request](https://github.com/hyperdriven/hyperdrive/compare/)

At this point you're waiting on us. We will try to review your submissions within three days -- typically sooner than that. We may suggest some changes, improvements or alternative approaches -- please be open to feedback, as we have to consider how your change fits into the over all design and vision of the project.

Some things that will increase the chance that your pull request is accepted:

- Break unrelated changes up into multiple pull requests -- our change log is generated automatically from merged pull requests.
- Add your changes to one of the existing files in our package, unless you are introducing an entirely new concept (in which case, you can create a new file).
- Write tests. We use [testify](https://github.com/stretchr/testify) for our unit tests.
- Follow the Golang [style guide](https://github.com/golang/go/wiki/CodeReviewComments).
- Write a [good commit message](http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html).
