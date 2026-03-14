import React, { useState, useEffect } from "react";
import InteractionButton from "./Components/button"

export default function App() {
  const [stats, setStats] = useState({ likes: 0, dislikes: 0 });

  const fetchStats = async () => {
    try {
      const res = await fetch("http://localhost:8080/api/interactions");
      const data = await res.json();
      setStats(data);
    } catch (err) {
      console.error("API Error:", err);
    }
  };

  useEffect(() => {
    fetchStats();
  }, []);

  const handleAction = async (isLikeValue) => {
    await fetch("http://localhost:8080/api/interactions", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ is_like: isLikeValue }),
    });
    fetchStats();
  };

  const handleReset = async () => {
    await fetch("http://localhost:8080/api/reset", { method: "POST" });
    setStats({ likes: 0, dislikes: 0 });
  };

  return (
    <div className="min-h-screen bg-gray-100 flex flex-col items-center justify-center p-4">
      <div className="bg-white p-8 rounded-2xl shadow-xl max-w-md w-full text-center">
        <h1 className="text-3xl font-bold text-gray-800 mb-2">Feedback</h1>
        <p className="text-gray-500 mb-8">What do you think of this post?</p>
        
        <div className="flex justify-center gap-4 mb-10">
          <InteractionButton 
            isLike={true} 
            count={stats.likes} 
            onAction={handleAction} 
          />
          <InteractionButton 
            isLike={false} 
            count={stats.dislikes} 
            onAction={handleAction} 
          />
        </div>

        <button
          onClick={handleReset}
          className="text-sm text-gray-400 hover:text-red-500 transition-colors flex items-center justify-center gap-1 w-full"
        >
          <span className="text-lg">🔄</span> Reset All Data
        </button>
      </div>
    </div>
  );
}

