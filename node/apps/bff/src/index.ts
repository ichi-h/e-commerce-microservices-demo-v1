import { serve } from '@hono/node-server'
import { Hono } from 'hono'

import { createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { ProductService } from "@e-commerce/product-connect";

const transport = createConnectTransport({
  baseUrl: "http://localhost:1324",
});

const client = createPromiseClient(ProductService, transport);

const app = new Hono()

app.get('/', async (c) => {
  const res = await client.getProducts({
    limit: 1,
    offset: 0
  });
  return c.text(JSON.stringify(res, null, 2))
})

const port = 3030
console.log(`Server is running on port ${port}`)

serve({
  fetch: app.fetch,
  port
})
