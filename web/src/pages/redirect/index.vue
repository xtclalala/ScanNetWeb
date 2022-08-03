<script lang="ts">
export default {
  name: 'Y1tRedirect',
}
</script>
<script setup lang="ts">
import { unref } from 'vue'
import { useRouter } from 'vue-router'
const { currentRoute, replace } = useRouter()
const { params, query } = unref(currentRoute)
const { path, _redirect_type = 'path' } = params
const _path = Array.isArray(path) ? path.join('/') : path

if (_redirect_type === 'name') {
  replace({
    name: _path,
    query,
    params,
  })
} else {
  replace({
    path: _path.startsWith('/') ? _path : '/' + _path,
    query,
  })
}
</script>
