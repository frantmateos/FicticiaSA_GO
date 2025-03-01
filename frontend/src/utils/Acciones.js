import axios from 'axios';
const authToken = localStorage.getItem('token'); 
export async function login(userData) {
  console.log(userData)
  try {
    const response = await axios.post('http://localhost:8080/users/login', userData, {
      credentials: "include",
    });
    console.log('Login response: ', response);
    localStorage.setItem('token', response.data.Token);
    return response.data.Token;
  } catch (error) {
    console.error('Login error: ', error);
    throw error;
  }
}
 
export async function register(userData){
  try {
    const response = await axios.post('http://localhost:8080/users', userData);
    console.log('Register response:', response);
    return response.data;
  } catch (error) {
    console.error('Register error:', error);
    throw error;
  }
}

export async function insertUser({nombre, genero, atributos,maneja, lentes,diabetico, enfermedades }) {
  try {
      console.log("Enviando datos al servidor:", { nombre, genero, atributos,maneja, lentes,diabetico, enfermedades });
      const response = await axios.post('http://localhost:8080/users', 
          {nombre, genero, atributos,maneja, lentes,diabetico, enfermedades }, 
          {
              headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
          });
      console.log("Respuesta del servidor:", response);
      return response.data;
  } catch (error) {
      console.error('Error al crear users en Acciones.js:', error);
      throw error;
  }
}

export async function getUserById(userId) {
  try {
    console.log("este id estoy pasando: ", userId)
    const response = await axios.get(`http://localhost:8080/users/${userId}`, {
      headers: { 'Authorization': `Bearer ${authToken}` }
    });
    console.log('Hotel cargado:', response.data);
    return response.data;
  } catch (error) {
    console.error('Error al obtener los hoteles¡?:', error.response ? error.response.data : error.message);
    throw error;
  }
}

export async function updateUser(userId, { nombre, genero, atributos,maneja, lentes,diabetico, enfermedades, estado}) {
  try {
    console.log("este id estoy pasando: ", userId)
    const response = await axios.put(`http://localhost:8080/users`, {id: userId,nombre, genero, atributos,maneja, lentes,diabetico, enfermedades,estado }, {
      headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
    });
    return response.data;
  } catch (error) {
    console.error('Error al actualizar el hotel:', error);
    throw error;
  }
}

export async function getAllUsers() {
  try {
    const response = await axios.get('http://localhost:8080/users/all', {
      headers: { 'Authorization': `Bearer ${authToken}` }
    });
    console.log('users cargados:', response.data);
    return response.data;
  } catch (error) {
    console.error('Error al obtener los users:', error.response ? error.response.data : error.message);
    throw error;
  }
}


export async function tokenId(){
  const token = localStorage.getItem('token');
  if (!token) {
    throw new Error('No token found');
  }
  console.log("tokens: ",token);
  const val1 = await axios.get('http://localhost:8080/users/token', {
  headers: {
    'Authorization': `Bearer ${token}`
  }
});
const val2 = val1.data.idU
return val2
}

export async function tokenRole(){
const token = localStorage.getItem('token');
console.log("tokens: ",token);
const val1 = await axios.get('http://localhost:8080/users/token', {
headers: {
  'Authorization': token
}
});
console.log("val1: ",val1)
const val2 = val1.data.Adminu
console.log("val2: ",val2)
return val2
}