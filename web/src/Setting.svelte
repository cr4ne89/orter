<script>
  let itemName = '';
  let itemType = '';

  const addItem = async () => {
    try {
      const response = await fetch('http://localhost:8080/items', {
        method: 'POST',
        credentials: 'include',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: new URLSearchParams({ name: itemName, item_type: itemType }),
      });

      if (response.ok) {
        alert('Item added successfully');
        itemName = '';
        itemType = '';
      } else {
        alert('Error adding item, please try again.');
      }
    } catch (error) {
      console.error('Add item error:', error);
      alert('An error occurred while adding the item. Please try again later.');
    }
  };
</script>

<div class="settings-container">
  <h2>Settings</h2>
  <input type="text" bind:value={itemName} placeholder="Item Name" />
  <input type="text" bind:value={itemType} placeholder="Item Type (text/checkbox/time)" />
  <button on:click={addItem}>Add Item</button>
</div>

<style>
  .settings-container {
    max-width: 400px;
    margin: auto;
    text-align: center;
  }
  input {
    display: block;
    margin: 10px auto;
    padding: 8px;
    width: 80%;
  }
  button {
    padding: 10px 20px;
  }
</style>
