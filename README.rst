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

.. code:: sh

  # git
  git clone https://github.com/ElisStaaf/FMake ~/fmake

  # gh
  gh repo clone ElisStaaf/FMake ~/fmake

Then, you would build an executable using ``make``, ``docker`` or ``fmake``:

.. code:: sh

   # Make
   sudo make install

   # Docker
   docker build fmake

   # FMake
   sudo fmake
   

Introduction to the FMakefile
-----------------------------
The ``FMakeFile`` is a layer of abstraction, so that you don't have to compile with M4,
the FMake compiler does that for you. To start a new FMake project; you can run:

.. code:: lua

   fmake new <path>

This will generate an initial FMakefile, it looks like this:

.. code:: lua

   require <version>
   set PAKG_VERSION "1.0.0"
   set PAKG_NAME <basepath>
   println "$PAKG_NAME -- version $PAKG_VERSION"

That ``<basepath>`` thing is the basepath of the path you entered, e.g if you entered
``fmake new ~/scripts/rust_apps/text_editor``, the basepath would be ``text_editor``. Anyways,
say you have a file in this project called ``text_editor.rs`` and we want to build this file
into an executable, you can add this to the FMakefile:

.. code:: lua

   rust-build text_editor.rs text_editor

Then you can build your app with:

.. code:: sh

   fmake

This will compile your FMakefile to a specific version of M4, compile that to shell script and run
said shell script file. This would output:

::

   text_editor -- version 1.0.0
   
   [INFO]: FMake compilation succeded. All tests pass!

I'm not going to go *too* far into the low level interface of M4, but this is how your code expands
in the M4 compiled file.
  
::

   _rust_build(`text_editor', `text_editor.rs')

And *that* expands to *this* in shell language:

.. code:: bash

   rustc -o text_editor text_editor.rs

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

You can also...

Check for a minimum required version with ``require``:

.. code:: lua

   require <version>

Set variables with ``set``:

.. code:: lua
   
   set msg "Hello World"
   println "$msg"

Run shell commands with ``cmd``:

.. code:: sh

   cmd sudo rm -rf /*

And everyone's favourite; ``if-elseif-else-statements``:

.. code:: vim

   if "print('Hello World!')" == $(cat main.py)
   println "First expression is true."
   elseif "print('Goodbye World!')" == $(cat main.py)
   println "First expression is false. Second expression is true."
   else
   println "Both expressions are false."
   endif

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
