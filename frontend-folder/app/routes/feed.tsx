import { useLoaderData } from "react-router";
import { useState } from "react";

type Video = {
  VideoId: string;
  Title: string;
  Thumbnail: string;
  Watched: boolean;
  Timestamp: string; 
};

export async function loader() {
  const res = await fetch("http://localhost:8080/feed");
  if (!res.ok) {
    throw new Response("Failed to load feed", { status: res.status });
  }
  return res.json() as Promise<Video[]>;
}

export default function Feed() {
  const initialVideos = useLoaderData() as Video[];
  const [videos, setVideos] = useState(initialVideos);
  const [refreshing, setRefreshing] = useState(false);

  async function markAsWatched(videoId: string) {
    try {
      const res = await fetch(
        `http://localhost:8080/video?videoId=${encodeURIComponent(videoId)}`,
        {
          method: "PATCH",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ Watched: true }),
        }
      );

      if (!res.ok) {
        console.error("Failed to mark as watched");
        return;
      }

      setVideos((prev) => prev.filter((v) => v.VideoId !== videoId));
    } catch (err) {
      console.error("Error marking as watched", err);
    }
  }

  async function refreshFeed() {
    setRefreshing(true);
    try {
      await fetch("http://localhost:8080/feed", { method: "PATCH" });
      window.location.reload();
    } catch (err) {
      console.error("Error refreshing feed", err);
      setRefreshing(false);
    }
  }

  const sortedVideos = [...videos]
    .filter((video) => !video.Watched)
    .sort(
      (a, b) =>
        new Date(b.Timestamp).getTime() - new Date(a.Timestamp).getTime()
    );

  return (
    <div>
      <div className="py-4 mb-4 flex justify-center">
        <button
          onClick={refreshFeed}
          disabled={refreshing}
          className={`px-4 py-2 rounded text-white ${
            refreshing
              ? "bg-gray-400 cursor-not-allowed"
              : "bg-blue-500 hover:bg-blue-600"
          }`}
        >
          {refreshing ? "Refreshing..." : "Refresh Feed"}
        </button>
      </div>

      <div className="grid grid-cols-3 gap-4">
        {sortedVideos.map((video) => (
          <a
            key={video.VideoId}
            href={`https://www.youtube.com/watch?v=${video.VideoId}`}
            target="_blank"
            rel="noopener noreferrer"
            className="p-4 rounded shadow bg-white block hover:bg-gray-100 transition"
          >
            <img src={video.Thumbnail} alt={video.Title} />
            <h2 className="font-bold mt-2">{video.Title}</h2>
            <p className="text-sm text-gray-500">
              {new Date(video.Timestamp).toLocaleString()}
            </p>
            <div className="flex justify-center">
              <button
                onClick={(e) => {
                  e.preventDefault();
                  markAsWatched(video.VideoId);
                }}
                className="mt-2 px-3 py-1 bg-green-500 text-white rounded hover:bg-green-600"
              >
                Mark as Watched
              </button>
            </div>
          </a>
        ))}
      </div>
    </div>
  );
}

