kubectl autoscale deploy wordpress --cpu-percent=50 --min=1 --max=10 -n wordpress
kubectl autoscale deploy wordpress-mysql --cpu-percent=50 --min=1 --max=10 -n wordpress