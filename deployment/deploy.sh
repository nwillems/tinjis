#!/bin/bash -xe

# TODO: Consider something with namespaces, eg as a parameter

kubectl apply -f tinjis.yml
kubectl apply -f antaeus.yml
