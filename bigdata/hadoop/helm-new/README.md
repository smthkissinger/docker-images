# Hadoop Chart

[Hadoop](https://hadoop.apache.org/) is a framework for running large scale distributed applications.

This chart is primarily intended to be used for YARN and MapReduce job execution where HDFS is just used as a means to transport small artifacts within the framework and not for a distributed filesystem. Data should be read from cloud based datastores such as Google Cloud Storage, S3 or Swift.

## Chart Details

## Installing the Chart

To install the chart with the release name `hadoop` that utilizes 50% of the available node resources:

```
$ helm install --name hadoop $(stable/hadoop/tools/calc_resources.sh 50) stable/hadoop
```

> Note that you need at least 2GB of free memory per NodeManager pod, if your cluster isn't large enough, not all pods will be scheduled.

The optional [`calc_resources.sh`](./tools/calc_resources.sh) script is used as a convenience helper to set the `yarn.numNodes`, and `yarn.nodeManager.resources` appropriately to utilize all nodes in the Kubernetes cluster and a given percentage of their resources. For example, with a 3 node `n1-standard-4` GKE cluster and an argument of `50`, this would create 3 NodeManager pods claiming 2 cores and 7.5Gi of memory.

### Persistence

To install the chart with persistent volumes:

```
$ helm install --name hadoop $(stable/hadoop/tools/calc_resources.sh 50) \
  --set persistence.nameNode.enabled=true \
  --set persistence.nameNode.storageClass=standard \
  --set persistence.dataNode.enabled=true \
  --set persistence.dataNode.storageClass=standard \
  stable/hadoop
```

> Change the value of `storageClass` to match your volume driver. `standard` works for Google Container Engine clusters.

## Configuration

The following table lists the configurable parameters of the Hadoop chart and their default values.

| Parameter                                         | Description                                                                        | Default                                                          |
| ------------------------------------------------- | -------------------------------                                                    | ---------------------------------------------------------------- |
| `image`                                           | Hadoop image ([source](https://github.com/Comcast/kube-yarn/tree/master/image))    | `danisla/hadoop:{VERSION}`                                       |
| `imagePullPolicy`                                 | Pull policy for the images                                                         | `IfNotPresent`                                                   |
| `hadoopVersion`                                   | Version of hadoop libraries being used                                              | `{VERSION}`                                                      |
| `antiAffinity`                                    | Pod antiaffinity, `hard` or `soft`                                                 | `hard`                                                           |
| `hdfs.nameNode.pdbMinAvailable`                   | PDB for HDFS NameNode                                                              | `1`                                                              |
| `hdfs.nameNode.resources`                         | resources for the HDFS NameNode                                                    | `requests:memory=256Mi,cpu=10m,limits:memory=2048Mi,cpu=1000m`   |
| `hdfs.dataNode.replicas`                          | Number of HDFS DataNode replicas                                                   | `1`                                                              |
| `hdfs.dataNode.pdbMinAvailable`                   | PDB for HDFS DataNode                                                              | `1`                                                              |
| `hdfs.dataNode.resources`                         | resources for the HDFS DataNode                                                    | `requests:memory=256Mi,cpu=10m,limits:memory=2048Mi,cpu=1000m`   |
| `yarn.resourceManager.pdbMinAvailable`            | PDB for the YARN ResourceManager                                                   | `1`                                                              |
| `yarn.resourceManager.resources`                  | resources for the YARN ResourceManager                                             | `requests:memory=256Mi,cpu=10m,limits:memory=2048Mi,cpu=1000m`   |
| `yarn.nodeManager.pdbMinAvailable`                | PDB for the YARN NodeManager                                                       | `1`                                                              |
| `yarn.nodeManager.replicas`                       | Number of YARN NodeManager replicas                                                | `2`                                                              |
| `yarn.nodeManager.parallelCreate`                 | Create all nodeManager statefulset pods in parallel (K8S 1.7+)                     | `false`                                                          |
| `yarn.nodeManager.resources`                      | Resource limits and requests for YARN NodeManager pods                             | `requests:memory=2048Mi,cpu=1000m,limits:memory=2048Mi,cpu=1000m`|
| `persistence.nameNode.enabled`                    | Enable/disable persistent volume                                                   | `false`                                                          | 
| `persistence.nameNode.storageClass`               | Name of the StorageClass to use per your volume provider                           | `-`                                                              |
| `persistence.nameNode.accessMode`                 | Access mode for the volume                                                         | `ReadWriteOnce`                                                  |
| `persistence.nameNode.size`                       | Size of the volume                                                                 | `50Gi`                                                           |
| `persistence.dataNode.enabled`                    | Enable/disable persistent volume                                                   | `false`                                                          | 
| `persistence.dataNode.storageClass`               | Name of the StorageClass to use per your volume provider                           | `-`                                                              |
| `persistence.dataNode.accessMode`                 | Access mode for the volume                                                         | `ReadWriteOnce`                                                  |
| `persistence.dataNode.size`                       | Size of the volume                                                                 | `200Gi`                                                          |

## Related charts

The [Zeppelin Notebook](https://github.com/kubernetes/charts/tree/master/stable/zeppelin) chart can use the hadoop config for the hadoop cluster and use the YARN executor:

```
helm install --set hadoop.useConfigMap=true stable/zeppelin
```

# References

- Original K8S Hadoop adaptation this chart was derived from: https://github.com/Comcast/kube-yarn

修改宿主机resolv.conf
search ceph.svc.cluster.local default.svc.cluster.local svc.cluster.local cluster.local
nameserver 10.233.0.3
nameserver 8.8.8.8
options timeout:2 attempts:3 rotate single-request-reopenbel node <nodename> node-type=storage

给存储节点打上标签(必须)
所有的存储节点
kubectl label node <nodename> hdfs=enabled
或者所有节点都打上标签
kubectl label nodes hdfs=enabled --all

修改记录
1.hdfs datanode使用hostnetwork
2.hdfs datanode使用daemonset
3.hdfs datanode存储使用hostdir /hdfs/<集群名字>/datanode/
4.hdfs namenode使用hostnetwork
5.hdfs namenode使用statefulset
6.所有hdfs使用指定节点 hdfs: enabled
7.自动设置HADOOP_HEAPSIZE