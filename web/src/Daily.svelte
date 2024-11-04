<script>
  export let date;
  let reflection = '';

  const submit = async () => {
    try {
      const response = await fetch('http://localhost:8080/daily-log', {
        method: 'POST',
        credentials: 'same-origin',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: new URLSearchParams({ date, items: reflection }),
      });

      if (response.ok) {
        alert('Daily reflection saved successfully');
        reflection = '';
      } else {
        alert('Error saving daily reflection, please try again.');
      }
    } catch (error) {
      console.error('Submit reflection error:', error);
      alert('An error occurred while saving your reflection. Please try again later.');
    }
  };
</script>

<div class="daily-container">
  <h2>Daily Reflection for {date}</h2>
  <textarea bind:value={reflection} placeholder="Write your reflection here..."></textarea>
  <button on:click={submit}>Submit</button>
</div>

<style>
  .daily-container {
    max-width: 400px;
    margin: auto;
    text-align: center;
  }
  textarea {
    width: 100%;
    height: 100px;
    margin: 10px 0;
    padding: 8px;
  }
  button {
    padding: 10px 20px;
  }
</style>
