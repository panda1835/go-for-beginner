# 1. History of Go

The idea of Go programming language was conceived by 3 Google-ers Robert Griesemer, Rob Pike, and Ken Thompson in 2007. Stories rumored that at the time, the three were so tired of the C++'s compilation and execution complexity at Google that they wanted to create a new language which upheld the "best" characteristics of widely used languages at the time. The new language should inherit the runtime efficency of C/C++, but at the same time maintained readability and usability of Python and Javascript. Moreover, it should also be able to adapt to the new era of multicore computing and networking. Those were the main initial ideas behind the Go language (also known as Golang because its former domain name was golang.org). In November 10, 2009, the authors published it as an open-source project and accepted contributions globally. And in March 2012, they release Go 1.0.

Go was born to solve many problems, among which 2 were the most notable: the development speed and concurrency. In terms of development speed, the Go's creators tried to bridge the gap between the quickness of code developement and program performance. Notably, Go is a "weird" combination of both static and dynamic typed languages, which makes it able to offer the comfort of interpreted languages such as Python, and the efficency of compiled language such as C/C++ or Java. On the source code compilation, Go offers high-performance compiling mechanism that can helps compile a huge system in just a range of seconds. For concurrency, Go has its own mechanism of concurrency that uses less memory compared to Threads from other languages. Besides, it also has a particular design pattern to channel data among running threads that keeps the data safe from concurrent modification, thus reduce the data loss burden.

Since the moment it was conceived at Google, Go was used and backed up a lot by its motherland. It was trusted by developers of Google's index service, which was the backbone of the whole Google searching system [\[1\]][1]. In another department, it is being used to optimize page load for Chrome Optimization Guide [\[2\]][2].

Go is a general-purpose language that is suitable for scaling up large system because of its capability at parallel programming and network concurrecy. As a result, besides Google, other from-big-to-small tech companies also adopts Go to scale up their systems and infrastructures. 
Dropbox shares their story of replacing Python, their time-honoured programming language, to Go to build a more scalable infrastructure to serve a growing community of users [\[3\]][3]. Although Go is still a young player, it seems like Go is a naturally fit for their need that they even write and open-source many standardized Go libraries to the community. 
In Microsoft, Go is the first-class citizen of their Azure infrastructure. Because of its high speed and ability to compile static non-dependency binary files, it attracts huge investement from the Microsoft Azure development team to provide best experiences for Go developers [\[4\]][4]. Other than that, Microsoft also invests in Go by developing Go-extension in Visual Studio Code (VSC) that makes experiences with Go in VSC a great experience.

There are a lot more stories of big tech companies sharing about their adoptation with Go and why/how they choose Go for their purposes. More stories can be found here [\[5\]][5].

Despite its presence in many modern cloud infrastructures and web services, Go is a bad option for client-side applications (such as GUI applications). Nevertheless, it is a great choice for products that are server-centered such as websites, network services, APIs, etc.



## References
1. [How Google's Core Data Solutions Team Uses Go][1]
2. [Chrome Content Optimization Service Runs on Go][2]
3. [Open Sourcing Our Go Libraries][3]
4. [Ready, Set, Go lang: Brian Ketelsen explains Go’s fast growing popularity][4]
5. [The Go website: Case Studies][5]
6. [Geeks for Geeks: Go Programming Language (Introduction)][6]
7. [A short history of the Go language][7]
8. Go in Action: Chapter 01 - Introducing Go


[1]: https://go.dev/solutions/google/coredata
[2]: https://go.dev/solutions/google/chrome
[3]: https://dropbox.tech/infrastructure/open-sourcing-our-go-libraries
[4]: https://cloudblogs.microsoft.com/opensource/2018/02/21/go-lang-brian-ketelsen-explains-fast-growth/
[5]: https://go.dev/solutions/#case-studies
[6]: https://www.geeksforgeeks.org/go-programming-language-introduction/
[7]: https://codilime.com/blog/go-programming-language-everything-you-should-know/