# training-docker

## Vagrant

* https://www.virtualbox.org/wiki/Downloads
* https://www.vagrantup.com/

```console
$ vagrant up
$ vagrant ssh
vagrant@ubuntu-bionic:~$ cd /vagrant
```

## multi stage

```console
vagrant@ubuntu-bionic:/vagrant/multi_stage$ docker build -t multistage:latest .
vagrant@ubuntu-bionic:/vagrant/multi_stage$ docker run -p 8000:8000 -d multistage:latest
vagrant@ubuntu-bionic:/vagrant/multi_stage$ docker ps
vagrant@ubuntu-bionic:/vagrant/multi_stage$ curl localhost:8000
```

## layer

```console
vagrant@ubuntu-bionic:/vagrant/layer$ docker build -t layer:latest
vagrant@ubuntu-bionic:/vagrant/layer$ docker save layer:latest > layer.tar
vagrant@ubuntu-bionic:/vagrant/layer$ mkdir -p ./tmp
vagrant@ubuntu-bionic:/vagrant/layer$ tar xvf layer.tar -C ./tmp
vagrant@ubuntu-bionic:/vagrant/layer$ tree tmp/
tmp/
├── 3855714064f0be0923497b7c97a39b60772f435150ffed09e2b951e852f955f1
│   ├── VERSION
│   ├── json
│   └── layer.tar
├── 6244fb2f12919689e8bb521f2ea97b642a1d832b34a43d34e419ad15dbae6603
│   ├── VERSION
│   ├── json
│   └── layer.tar
├── 9eb248cd077fb9f41f2cae8c184b4ce7f919583629d648dd69d6809c50d13c7c.json
├── c0c96aecb724fb8b18fcfc9f880ca6242a0f4f8f57ded2590cf047007dbef999
│   ├── VERSION
│   ├── json
│   └── layer.tar
├── manifest.json
└── repositories

```
