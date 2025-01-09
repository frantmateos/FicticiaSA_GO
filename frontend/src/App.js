import React, { useEffect } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import LoginRegister from './components/LoginRegister/LoginRegister'; 
import Home from './components/home/home';
import './App.css';

function App() {
  
  return (
    <BrowserRouter>
      <div className="App">
        <Routes>
       
          <Route path="/" element={<LoginRegister />} />

         
          <Route path="/home" element={<Home />} />


      

        </Routes>
      </div>
    </BrowserRouter>
  );
  
}

export default App;
