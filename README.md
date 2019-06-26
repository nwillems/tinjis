
Hi there

Instructions:
1. Build and docker-composed the antaeus service
2. Created own service using Golang - see `./payment`
3. Done and done. Except - figured that antaeus will only continue processing
   invoices if the result of the operation was true. Eg. The tes should then be
   something along the lines of calling the endpoint N-times, where
   N=#invoices.
4. Done :-) And tested locally on minikube. See `./deployment`
5. See below sections :-)


## New deployment
The implemented setup suffers from a lot of duplication. This could somewhat be
alleviated by using helm-charts. I've never worked with kotlin before, but
having seen and interacted with it during this project, I feel its quite a nice
language.

Another area where making the deployment into helm-chart(s), would be
configuration values that could differ between eg staging and prod.

To ensure that the two services stay independent, I would make sure to use a
contract or mock service. Eg pact. If it is the same team responsible for both
services, I guess it wouldn't matter too much, but it could help. If it's a
cross team dependency, I would definately implement something like contract
testing.

The payment service famously lacks testing, and a propper CI setup, this could
be achieved using something like travis/... which I would prefer over Jenkins.
1) I/We don't host it 2) Jenkins and plugins and configuring it :-(
Unfortunately I have great experience with Jenkins :-)

## Limiting access to deployments
I have figured two ways I like, but can't decide which I like better.

First, namespaces and RBAC. I'm not too familiar with RBAC, but I figure it
could probably allow us to limit who deploys where. This would be a great
solution, as it is already made.

Second, disallow deployments altogether directly to the cluster, but implement
a middle-ware-sort-of-thing, that will enfore our policies. This could also be
used to implement such a thing as when you are allowed to deploy, and if you
are adhering to SLOs or...
Downside to this, maybe it's already made, but it seems to be a new component.

Weighing the tradeoffs, I guess initially I would go with namespaces and RBAC,
and then when the need arises, investigate something more advanced. Maybe istio
can do something?

## Limiting service-access
There could be many layers where this could be implemented. So it sort of
depends on the level of separation we would like.

If we wan't total separation, eg not even a TCP-SYN thing, then I would send
the traffic outside the cluster and limit it through other means. This would
also allow to separate the two into separate clusters.

My best solution though, would be to look more into SPIFFE, which could provide
us with service identities, which we could then use to implement the right
access limitations. This solutions balances between application layer and
lower-level network layer.

## Comments on the current setup
The interaction between antaeus and the payment service should probably do
something about retires, in case the service is failing hard/soft.

The current antaeus implementation, where it stops processing payments as soon
as one fails, seems also related to the above comment. 

* The readinessProbe for antaeus has been set extremely high, this is due to
  poor performance on my machine. I timed the startup of the container and it
  took 13min20secs. I noted during this, that it downloads gradle everytime it
  starts, wondering how one would ensure its moved from the build container to
  the prod one.
