function toggleCell(cell) {
    cell.classList.toggle('selected');
    
    // Vérifier si on a un bingo
    checkBingo();
}

function checkBingo() {
    // Implémentez la logique de vérification de bingo ici
    // Par exemple, vérifier une ligne, colonne ou diagonale complète
    console.log("Vérification du bingo...");
}
