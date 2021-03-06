<!-- BEGIN MUNGE: UNVERSIONED_WARNING -->

<!-- BEGIN STRIP_FOR_RELEASE -->

<img src="http://kubernetes.io/img/warning.png" alt="WARNING"
     width="25" height="25">
<img src="http://kubernetes.io/img/warning.png" alt="WARNING"
     width="25" height="25">
<img src="http://kubernetes.io/img/warning.png" alt="WARNING"
     width="25" height="25">
<img src="http://kubernetes.io/img/warning.png" alt="WARNING"
     width="25" height="25">
<img src="http://kubernetes.io/img/warning.png" alt="WARNING"
     width="25" height="25">

<h2>PLEASE NOTE: This document applies to the HEAD of the source tree</h2>

If you are using a released version of Kubernetes, you should
refer to the docs that go with that version.

<!-- TAG RELEASE_LINK, added by the munger automatically -->
<strong>
The latest release of this document can be found
[here](http://releases.k8s.io/release-1.1/docs/user-guide/kubectl/kubectl_edit.md).

Documentation for other releases can be found at
[releases.k8s.io](http://releases.k8s.io).
</strong>
--

<!-- END STRIP_FOR_RELEASE -->

<!-- END MUNGE: UNVERSIONED_WARNING -->

## kubectl edit

Edit a resource on the server

### Synopsis


Edit a resource from the default editor.

The edit command allows you to directly edit any API resource you can retrieve via the
command line tools. It will open the editor defined by your KUBE_EDITOR, or EDITOR
environment variables, or fall back to 'vi' for Linux or 'notepad' for Windows.
You can edit multiple objects, although changes are applied one at a time. The command
accepts filenames as well as command line arguments, although the files you point to must
be previously saved versions of resources.

The files to edit will be output in the default API version, or a version specified
by --output-version. The default format is YAML - if you would like to edit in JSON
pass -o json. The flag --windows-line-endings can be used to force Windows line endings,
otherwise the default for your operating system will be used.

In the event an error occurs while updating, a temporary file will be created on disk
that contains your unapplied changes. The most common error when updating a resource
is another editor changing the resource on the server. When this occurs, you will have
to apply your changes to the newer version of the resource, or update your temporary
saved copy to include the latest resource version.

```
kubectl edit (RESOURCE/NAME | -f FILENAME)
```

### Examples

```
  # Edit the service named 'docker-registry':
  $ kubectl edit svc/docker-registry

  # Use an alternative editor
  $ KUBE_EDITOR="nano" kubectl edit svc/docker-registry

  # Edit the service 'docker-registry' in JSON using the v1 API format:
  $ kubectl edit svc/docker-registry --output-version=v1 -o json
```

### Options

```
  -f, --filename=[]: Filename, directory, or URL to file to use to edit the resource
  -o, --output="yaml": Output format. One of: yaml|json.
      --output-version="": Output the formatted object with the given version (default api-version).
      --save-config[=false]: If true, the configuration of current object will be saved in its annotation. This is useful when you want to perform kubectl apply on this object in the future.
      --windows-line-endings[=false]: Use Windows line-endings (default Unix line-endings)
```

### Options inherited from parent commands

```
      --alsologtostderr[=false]: log to standard error as well as files
      --api-version="": The API version to use when talking to the server
      --certificate-authority="": Path to a cert. file for the certificate authority.
      --client-certificate="": Path to a client certificate file for TLS.
      --client-key="": Path to a client key file for TLS.
      --cluster="": The name of the kubeconfig cluster to use
      --context="": The name of the kubeconfig context to use
      --insecure-skip-tls-verify[=false]: If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure.
      --kubeconfig="": Path to the kubeconfig file to use for CLI requests.
      --log-backtrace-at=:0: when logging hits line file:N, emit a stack trace
      --log-dir="": If non-empty, write log files in this directory
      --log-flush-frequency=5s: Maximum number of seconds between log flushes
      --logtostderr[=true]: log to standard error instead of files
      --match-server-version[=false]: Require server version to match client version
      --namespace="": If present, the namespace scope for this CLI request.
      --password="": Password for basic authentication to the API server.
  -s, --server="": The address and port of the Kubernetes API server
      --stderrthreshold=2: logs at or above this threshold go to stderr
      --token="": Bearer token for authentication to the API server.
      --user="": The name of the kubeconfig user to use
      --username="": Username for basic authentication to the API server.
      --v=0: log level for V logs
      --vmodule=: comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO

* [kubectl](kubectl.md)	 - kubectl controls the Kubernetes cluster manager

###### Auto generated by spf13/cobra on 8-Dec-2015

<!-- BEGIN MUNGE: GENERATED_ANALYTICS -->
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/docs/user-guide/kubectl/kubectl_edit.md?pixel)]()
<!-- END MUNGE: GENERATED_ANALYTICS -->
