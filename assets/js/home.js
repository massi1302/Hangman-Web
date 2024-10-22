document.addEventListener('DOMContentLoaded', () => {
    // Sélectionner tous les boutons de difficulté
    const difficultyButtons = document.querySelectorAll('.difficulty-buttons .button');
    
    // Ajouter un écouteur d'événement pour chaque bouton
    difficultyButtons.forEach(button => {
        button.addEventListener('click', () => {
            // Retirer la classe 'selected' de tous les boutons
            difficultyButtons.forEach(btn => btn.classList.remove('selected'));
            // Ajouter la classe 'selected' au bouton cliqué
            button.classList.add('selected');
        });
    });
});