# Read Chart.yaml from Helm Charts

You can read Chart.yaml from your Helm Charts and use them as outputs:

```yaml
  - name: Read Helm Chart
    id: chart-info
    uses: miraai/read-helm-chart-yaml@v0.1
    with:
        path: tests/example
```