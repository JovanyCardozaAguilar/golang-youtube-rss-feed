import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
  index("routes/feed.tsx"),
  route("channels", "routes/channels.tsx"),
  route("addChannel", "routes/addChannel.tsx"),
  route("watched", "routes/watched.tsx"),
] satisfies RouteConfig;
