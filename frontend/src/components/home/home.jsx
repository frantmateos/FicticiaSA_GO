import React, { useState, useEffect } from 'react';
import { insertUser, updateUser, getAllUsers } from '../../utils/Acciones.js';
import './home.css';
import { useNavigate } from 'react-router-dom';
import { tokenRole } from '../../utils/Acciones';

const MisUsuarios = () => {
  const [showAddDialog, setShowAddDialog] = useState(false);
  const [showEditDialog, setShowEditDialog] = useState(false);
  const [users, setUsers] = useState([]);
  const [mensaje, setMensaje] = useState('');
  const [filteredUsers, setFilteredUsers] = useState([]);
  const [isAdmin, setRole] = useState('');
  const [errors, setErrors] = useState({});

  const navigate = useNavigate();

  const [selectedUser, setSelectedUser] = useState(null);

  const [nombre, setNombre] = useState('');
  const [genero, setGenero] = useState('');
  const [atributos, setAtributos] = useState('');
  const [maneja, setManeja] = useState(false);
  const [lentes, setLentes] = useState(false);
  const [diabetico, setDiabetico] = useState(false);
  const [enfermedades, setEnfermedades] = useState('');
  const [estado, setEstado] = useState(false);

  const openAddDialog = () => {
    setShowAddDialog(true);
    document.body.style.overflow = 'hidden';
  };

  const closeAddDialog = () => {
    setShowAddDialog(false);
    document.body.style.overflow = 'auto';
    resetForm();
  };

  const openEditDialog = (user) => {
    setSelectedUser(user);
    setNombre(user.nombre || '');
    setGenero(user.genero || '');
    setAtributos(user.atributos || '');
    setManeja(user.maneja || false);
    setLentes(user.lentes || false);
    setDiabetico(user.diabetico || false);
    setEnfermedades(user.enfermedades || '');
    setEstado(user.estado || false);
    setErrors({}); 
    setShowEditDialog(true);
  };

  const closeEditDialog = () => {
    setShowEditDialog(false);
    setSelectedUser(null);
    resetForm();
  };

  const resetForm = () => {
    setNombre('');
    setGenero('');
    setAtributos('');
    setManeja(false);
    setLentes(false);
    setDiabetico(false);
    setEnfermedades('');
    setEstado(false);
    setErrors({});
  };

  useEffect(() => {
    const fetchRole = async () => {
      try {
        const role = await tokenRole();
        setRole(role);
        console.log("role este: ",role)
      } catch (error) {
        console.error('Error fetching role:', error);
      }
    };
    fetchRole();
  }, []);

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const usersData = await getAllUsers();
        setUsers(usersData);
        setFilteredUsers(usersData);
      } catch (error) {
        console.error('Error fetching users:', error);
      }
    };
    fetchUsers();
  }, []);

  const validateFields = () => {
    const newErrors = {};
    if (!nombre || nombre.length < 5) {
      newErrors.nombre = 'El nombre debe tener al menos 5 caracteres.';
    }
    if (!genero) {
      newErrors.genero = 'El género es obligatorio.';
    }
    if (!atributos) {
      newErrors.atributos = 'Los atributos son obligatorios.';
    }
    setErrors(newErrors);
    return Object.keys(newErrors).length === 0; 
  };

  const handleInsertUser = async (e) => {
    e.preventDefault();
    if (!validateFields()) return;

    const userData = { nombre, genero, atributos, maneja, lentes, diabetico, enfermedades, estado };
    try {
      const newUser = await insertUser(userData);
      setUsers((prevUsers) => [...prevUsers, newUser]);
      setMensaje('Usuario creado exitosamente.');
      closeAddDialog();
    } catch (error) {
      setMensaje('Error al crear usuario');
      console.error('Error en handleInsertUser:', error.response ? error.response.data : error.message);
    }
  };

  const handleUpdateUser = async (e) => {
    e.preventDefault();
    if (!validateFields()) return;
    const userData = {
      id: selectedUser.id,
      nombre,
      genero,
      atributos,
      maneja,
      lentes,
      diabetico,
      enfermedades,
      estado,
    };

    try {
      const updatedUser = await updateUser(selectedUser.id, userData);
      setUsers((prevUsers) =>
        prevUsers.map((user) => (user.id === updatedUser.id ? updatedUser : user))
      );
      setMensaje('Usuario actualizado exitosamente.');
      closeEditDialog();
    } catch (error) {
      setMensaje('Error al actualizar usuario');
      console.error('Error en handleUpdateUser:', error.response ? error.response.data : error.message);
    }
  };

  return (
    <div className="contenedor-reserva">
      <h1>Usuarios</h1>
      <button className="Agregar-Usuario" onClick={openAddDialog}>
        Agregar Usuario
      </button>

      <ul className="Lista-usuarios">
        {filteredUsers.length > 0 ? (
          filteredUsers.filter((data) => !data.admin).map((data) => (
            <li key={data.id} className="bloque">
              <h2>{data.nombre}</h2>
              <p>{data.atributos || 'Sin atributos adicionales.'}</p>
              <p>Estado: {data.estado ? 'Activo' : 'Inactivo'}</p>
              <button className="boton-editar" onClick={() => openEditDialog(data)}>
                Actualizar
              </button>
            </li>
          ))
        ) : (
          <p>No se encontraron usuarios.</p>
        )}
      </ul>

      {showAddDialog && (
        <div className="modal">
          <form onSubmit={handleInsertUser}>
            <div className="modal-content">
              <h2>Agregar Nuevo Usuario</h2>
              <input
                type="text"
                value={nombre}
                onChange={(e) => setNombre(e.target.value)}
                placeholder="Nombre del Usuario"
              />
              {errors.nombre && <p style={{ color: 'red' }}>{errors.nombre}</p>}

              <input
                type="text"
                value={genero}
                onChange={(e) => setGenero(e.target.value)}
                placeholder="Género"
              />
              {errors.genero && <p style={{ color: 'red' }}>{errors.genero}</p>}

              <textarea
                value={atributos}
                onChange={(e) => setAtributos(e.target.value)}
                placeholder="Atributos"
              />
              {errors.atributos && <p style={{ color: 'red' }}>{errors.atributos}</p>}

              <label>
                <input
                  type="checkbox"
                  checked={maneja}
                  onChange={(e) => setManeja(e.target.checked)}
                />
                Maneja
              </label>
              <label>
                <input
                  type="checkbox"
                  checked={lentes}
                  onChange={(e) => setLentes(e.target.checked)}
                />
                Usa lentes
              </label>
              <label>
                <input
                  type="checkbox"
                  checked={diabetico}
                  onChange={(e) => setDiabetico(e.target.checked)}
                />
                Diabético
              </label>
              <input
                type="text"
                value={enfermedades}
                onChange={(e) => setEnfermedades(e.target.value)}
                placeholder="Enfermedades"
              />
                

              <button type="submit">Agregar</button>
              <button type="button" onClick={closeAddDialog}>Cancelar</button>
            </div>
          </form>
        </div>
      )}

      {showEditDialog && (
        <div className="modal">
          <form onSubmit={handleUpdateUser}>
            <div className="modal-content">
              <h2>Actualizar Usuario</h2>
              <input
                type="text"
                value={nombre}
                onChange={(e) => setNombre(e.target.value)}
                placeholder="Nombre del Usuario"
              />
              {errors.nombre && <p style={{ color: 'red' }}>{errors.nombre}</p>}

              <input
                type="text"
                value={genero}
                onChange={(e) => setGenero(e.target.value)}
                placeholder="Género"
              />
              {errors.genero && <p style={{ color: 'red' }}>{errors.genero}</p>}

              <textarea
                value={atributos}
                onChange={(e) => setAtributos(e.target.value)}
                placeholder="Atributos"
              />
              {errors.atributos && <p style={{ color: 'red' }}>{errors.atributos}</p>}
              <label>
                <input
                  type="checkbox"
                  checked={maneja}
                  onChange={(e) => setManeja(e.target.checked)}
                />
                Maneja
              </label>
              <label>
                <input
                  type="checkbox"
                  checked={lentes}
                  onChange={(e) => setLentes(e.target.checked)}
                />
                Usa lentes
              </label>
              <label>
                <input
                  type="checkbox"
                  checked={diabetico}
                  onChange={(e) => setDiabetico(e.target.checked)}
                />
                Diabético
              </label>
              <input
                type="text"
                value={enfermedades}
                onChange={(e) => setEnfermedades(e.target.value)}
                placeholder="Enfermedades"
              />
              <label>
                <input
                  type="checkbox"
                  checked={estado}
                  onChange={(e) => setEstado(e.target.checked)}
                />
                Estado Activo
              </label>

              <button type="submit">Actualizar</button>
              <button type="button" onClick={closeEditDialog}>Cancelar</button>
            </div>
          </form>
        </div>
      )}
    </div>
  );
};

export default MisUsuarios;
