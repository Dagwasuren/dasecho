# Contribute


## GitHub and Go: forking, pull requests, and go-getting

During the Go bootcamp that took place in Los Angeles some days ago (golangbootcamp.com) someone asked me a very interesting question.

Question: how do you handle pull requests from different GitHub repositories?

The question seemed pretty simple to answer: you fork the project, then you do "go get" of the fork.

That answer is simple, but incorrect! If you know why ... well, then just go to the answer below.

It is incorrect, because if you fork a repository that has import statements referring to itself, they will be broken in the fork.

Before it gets more complicated let's make this a bit more concrete. You want to send a pull request to fix my awesome project cooltool in github.com/campoy/cooltool.


First step is to fork the repository, which creates github.com/you/cooltool. Easy!

Then you download the source code of the fork using "go get", this will put all the code from the repository in $GOPATH/src/github.com/you/cooltool.

`
$ go get github.com/you/cooltool
`


Interestingly enough that also downloads github.com/campoy/cooltool. Why?

Well, let's see the contents of github.com/you/cooltool/main.go.

Do you see that? The import path refers to github.com/campoy/cooltol!

You could change it to refer to github.com/you/cooltool, but then you would run into conflicts when you send the pull request.

Answer: forks and remotes

There's a very simple solution to this problem consisting in three simple steps.

First fork the repository, as before this will create a github.com/you/cooltool.

Then let's get the original code using go get. Yes, the original instead of the fork.


$ go get github.com/campoy/cooltool

$ git remote add fork https://github.com/you/cooltool.git


This added your fork as a remote, which means that once you've modified some files and committed the changes you can now run:


$ git push fork

Conclusion

GitHub is a great place to have your code and ,once you know how, sending pull requests for Go repositories is trivial!
