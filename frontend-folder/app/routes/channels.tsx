import { useLoaderData } from "react-router";
import { useState } from "react";

type Channel = {
  ChannelId: string;
  Handle: string;
  Username: string;
  Avatar: string;
};

export async function loader() {
  const res = await fetch("http://localhost:8080/channelFeed");
  if (!res.ok) {
    throw new Response("Failed to load channels", { status: res.status });
  }
  return res.json() as Promise<Channel[]>;
}

export default function Channels() {
  const initialChannels = useLoaderData() as Channel[];
  const [channels, setChannels] = useState(initialChannels);

  async function deleteChannel(ChannelId: string) {
    const res = await fetch(`http://localhost:8080/channel?channelId=${encodeURIComponent(ChannelId)}`, {
      method: "DELETE",
    });

    if (res.ok) {
      setChannels(channels.filter((ch) => ch.ChannelId !== ChannelId));
    } else {
      console.error("Failed to delete channel");
    }
  }

  return (
    <div className="max-w-3xl mx-auto">
      <h1 className="text-2xl font-bold mb-4">Channels</h1>
      <ul className="space-y-2">
        {channels.map((ch) => (
          <li
            key={ch.ChannelId}
            className="p-4 bg-white rounded shadow flex items-center justify-between"
          >
            <div className="flex items-center">
              <img
                src={ch.Avatar}
                alt={ch.Username}
                className="w-12 h-12 rounded-full mr-4"
              />
              <div>
                <div className="text-black font-bold">{ch.Username}</div>
                <div className="text-gray-600">@{ch.Handle}</div>
              </div>
            </div>
            <button
              onClick={() => deleteChannel(ch.ChannelId)}
              className="bg-red-500 text-white px-3 py-1 rounded hover:bg-red-600 transition"
            >
              Delete
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
}

