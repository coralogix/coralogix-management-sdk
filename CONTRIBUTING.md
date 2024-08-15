# How To Contribute? 

We are excited that you want to participate in this repository, there are two general ways you can do that:

1. Opening an issue
2. Opening a pull request

Each of those are valuable for improving what we do here. Let's look at more details. 

## Opening An Issue

By pointing out bugs, issues, or even requesting features you can impact where the future development is going without writing any code. However we do need a detailed explanation and supporting material within the issue, so here is a list of things that have to be covered:

1. The goal: what are you trying to achieve
2. The problem: what issues are you facing 
3. (Pseudo) code for the reproducing the issue 
    - Base the code on one of the examples
    - Pseudo code only for new features to show the ideal outcome
4. Output of the code, expected and actual

These aspects allow us to follow up on the issue quicker and plan the design and implementation. If you want to participate in the coding itself, keep reading.

## Opening A Pull Request

A pull request (PR) is a great way of fixing issues, small and large, as well as adding new features. An ideal PR refers to the issue it is solving, contains tests, documentation, and a sensible amount of code. In a list, the following formal requirements have to be met:

1. An issue describing the problem and sketching out the solution
    - For code contributions only, docs or tests can do without
2. The contribution passes CI
3. The CLA has been signed
4. Code contributions are tested and documented as appropriate
5. Large contributions should be split into multiple stand-alone PRs

With these formal requirements fulfilled, a reviewer will have an easy time accepting the code and we can keep the quality of the repository at a reasonable level. For technical details on how to run the tests/create docs/... please refer to the [README.md]() of this repository.

### Git pre-commit hook
We provide a git pre-commit hook that you can use in order to verify whether your code satisfies some basic quality checks. You can install it by simply running the `install_hooks.sh` script in the root of the repository. You're not forced to install it, but we highly recommend you do.



