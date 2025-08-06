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

export default function Watched() {
  const initialVideos = useLoaderData() as Video[];
  const [videos, setVideos] = useState(initialVideos);

  async function markAsUnwatched(videoId: string) {
    try {
      const res = await fetch(
        `http://localhost:8080/video?videoId=${encodeURIComponent(videoId)}`,
        {
          method: "PATCH",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ Watched: false }),
        }
      );

      if (!res.ok) {
        console.error("Failed to mark as unwatched");
        return;
      }

      setVideos((prev) => prev.filter((v) => v.VideoId !== videoId));
    } catch (err) {
      console.error("Error marking as unwatched", err);
    }
  }

  return (
    <div className="grid grid-cols-3 gap-4">
      {videos
        .filter((video) => video.Watched)
        .sort(
          (a, b) =>
            new Date(b.Timestamp).getTime() -
            new Date(a.Timestamp).getTime()
        )
        .map((video) => (
          <div
            key={video.VideoId}
            className="p-4 rounded shadow bg-white cursor-pointer"
            onClick={() =>
              window.open(
                `https://www.youtube.com/watch?v=${video.VideoId}`,
                "_blank"
              )
            }
          >
            <img src={video.Thumbnail} alt={video.Title} />
            <h2 className="font-bold mt-2">{video.Title}</h2>
            <p className="text-sm text-gray-500">
              {new Date(video.Timestamp).toLocaleString()}
            </p>
            <div className="flex justify-center">
              <button
                onClick={(e) => {
                  e.stopPropagation(); 
                  markAsUnwatched(video.VideoId);
                }}
                className="mt-2 px-3 py-1 bg-red-500 text-white rounded hover:bg-red-600"
              >
                Unwatch
              </button>
            </div>
          </div>
        ))}
    </div>
  );
}

