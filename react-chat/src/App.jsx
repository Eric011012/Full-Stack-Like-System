import React, { useState, useEffect } from 'react';

function App() {
  const [likes, setLikes] = useState(0);

  // 获取初始数据
  const fetchLikes = async () => {
    const res = await fetch('http://localhost:8080/api/like');
    const data = await res.json();
    setLikes(data.count);
  };

  useEffect(() => {
    fetchLikes();
  }, []);

  // 点赞逻辑
  const handleLike = async () => {
    await fetch('http://localhost:8080/api/like', { method: 'POST' });
    fetchLikes(); // 更新后重新获取最新数字
  };

  // 重置逻辑
  const handleReset = async () => {
    if (window.confirm("确定要清空所有点赞吗？")) {
      await fetch('http://localhost:8080/api/reset', { method: 'POST' });
      setLikes(0); // 前端立即归零
    }
  };

  return (
    <div style={{ textAlign: 'center', marginTop: '50px', fontFamily: 'Arial' }}>
      <h1>Likes: {likes}</h1>
      <button onClick={handleLike} style={buttonStyle}>👍 点赞</button>
      <br /><br />
      <button onClick={handleReset} style={{ ...buttonStyle, backgroundColor: '#ff4d4f' }}>
        🔄 重置归零
      </button>
    </div>
  );
}

const buttonStyle = {
  fontSize: '18px',
  padding: '10px 20px',
  cursor: 'pointer',
  borderRadius: '8px',
  border: 'none',
  backgroundColor: '#1890ff',
  color: 'white',
  margin: '5px'
};

export default App;