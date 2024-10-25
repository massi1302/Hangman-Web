document.addEventListener('DOMContentLoaded', () => {
   
    const difficultyButtons = document.querySelectorAll('.difficulty-buttons .button');
    
    
    difficultyButtons.forEach(button => {
        button.addEventListener('click', () => {
          
            difficultyButtons.forEach(btn => btn.classList.remove('selected'));
         
            button.classList.add('selected');
        });
    });
});