import React, { useState } from 'react';

function App() {
  const [name, setName] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault(); // Prevent the page from refreshing

    // Make a POST request to the backend with the form data
    fetch('http://localhost:8080/hello-world/submit', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ name }),
    })
      .then(response => response.text())
      .then(data => {
        console.log('Response from backend:', data);
      })
      .catch(error => {
        console.error('Error:', error);
      });
  };

  return (
    <div>
      <div className="min-h-screen flex items-center justify-center bg-blue-100">
        <h1 className="text-3xl font-bold text-blue-600">
          Hello, Tailwind CSS!
        </h1>
      </div>

      <div>
        <h1>Submit Your Details</h1>
        <form onSubmit={handleSubmit}>
          <div>
            <label>Name:</label>
            <input
              type="text"
              value={name}
              onChange={(e) => setName(e.target.value)}
            />
          </div>
          <button type="submit">Submit</button>
        </form>
      </div>
    </div>
    
  );
}

export default App;
