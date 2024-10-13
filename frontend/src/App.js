import React, { useState } from 'react';

function App() {
  const [habit, setHabit] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault(); // Prevent the page from refreshing

    // Make a POST request to the backend with the form data
    fetch(`${process.env.REACT_APP_BACKEND_URL}/hello-world/submit`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ habit }),
    })
      .then(response => response.text())
      .then(data => {
        console.log('Response from backend:', data);
        setHabit(''); // Clear the input after submission
      })
      .catch(error => {
        console.error('Error:', error);
      });
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
        <h2 className="text-2xl font-semibold text-center text-gray-800 mb-6">New Habit</h2>
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label className="block text-gray-600 font-medium mb-2" htmlFor="habit">
              Habit:
            </label>
            <input
              type="text"
              id="habit"
              value={habit}
              onChange={(e) => setHabit(e.target.value)}
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="Enter a new habit"
              required
            />
          </div>
          <button
            type="submit"
            className="w-full bg-blue-500 text-white font-semibold py-2 rounded-lg shadow-md hover:bg-blue-600 transition duration-200"
          >
            Add Habit
          </button>
        </form>
      </div>
    </div>
  );
}

export default App;
