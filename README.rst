FMake: Build software for idiots
================================
.. image:: https://img.shields.io/badge/Build%20(Fedora)-passing-2a7fd5?logo=fedora&logoColor=2a7fd5&style=for-the-badge
   :alt: Build = Passing
   :target: https://github.com/ElisStaaf/FMake
.. image:: https://img.shields.io/badge/Version-NET/1-38c747?style=for-the-badge
   :alt: Version = 1.0.0
   :target: https://github.com/ElisStaaf/FMake
.. image:: https://img.shields.io/badge/Language-Go-19cbe6?logo=go&style=for-the-badge
   :alt: Language = Python
   :target: https://github.com/ElisStaaf/FMake

FMake is build software focused on working. Unlike other build software this one doesn't work half of the
time. It also (right now at least) doesn't include anything special. You can only... Build stuff and print
stuff.

Requirements
------------
* `go`_
* `make`_ or `docker`_
* `m4`_
* `git`_ or `gh`_

Install
-------
To install, firstly, clone the git repo:

.. code:: bash

  # git
  git clone https://github.com/ElisStaaf/FMake ~/fmake

  # gh
  gh repo clone ElisStaaf/FMake ~/fmake

Then, you would build an executable using ``make`` or ``docker``:

.. code:: bash

   # Make
   sudo make install

   # Docker
   docker build fmake

Introduction to the FMakefile
-----------------------------
The ``FMakeFile`` is a layer of abstraction, so that you don't have to compile with M4, the FMake compiler
does that for you. Say you have a project with a file called ``main.rs``, you can create an ``FMakefile``
and write this into it:

.. code:: lua

   rust-build main.rs main

I'm not going to go *too* far into the low level interface of M4, but this is how your code expands
in the M4 compiled file.
  
::

   _rust_build(`main', `main.rs')

And *that* expands to *this* in shell language:

.. code:: bash

   rustc -o main main.rs

Comments in FMake start with ``--``:
  
.. code:: lua

  -- This is a comment, and it it awesome.

There are other compilers you can use in FMake, here's a showcase:
  
.. code:: lua

   -- This is the rust compiler, the one I showed earlier:
   rust-build main.rs main

   -- This is the GCC compiler:
   gcc-build main.c main

   -- This is the G++ compiler:
   g++-build main.cpp main

   -- And this is the Go compiler:
   go-build main.go main

``println`` statements also exist:

.. code:: lua

   println "Hello World!"

You can also invoke the compiler with many different flags, these are all of them (for now):

::

   -h, --help: Show help message.
   -v, --version: Print version name.
   -S: Save all tmp files.

.. _`go`: https://go.dev/doc/install
.. _`make`: https://www.gnu.org/software/make
.. _`docker`: https://docs.docker.com/engine/install/
.. _`m4`: https://www.linuxfromscratch.org/museum/lfs-museum/2.3.7/LFS-BOOK-2.3.7-HTML/x2018.html
.. _`git`: https://git-scm.com/downloads
.. _`gh`: https://github.com/cli/cli#installation
