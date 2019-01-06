package squareRoot

import (
    "github.com/acra5y/n-dilation-computer/eye"
    "gonum.org/v1/gonum/mat"
    "math"
)

// https://scholarworks.rit.edu/cgi/viewcontent.cgi?article=10419&context=theses

/*
    1. Declare some nonsingular matrix C with dimensions (n, n).
    2. Initialize i for number of iterations, S_0 = I and S_1 = C.
    3. Initialize Z = C − I.
    4. For i iterations or until S_i becomes too ill-conditioned, do S_{i+1} = 2S_i + (Z)(S_{i−1}),
    5. After iteration steps stop, find S_{i}^{−1}
    .
    6. Set n × n matrix Q = S_{i+1}(S_{i}^{−1}) − I.
*/

func inverseViaQR(m *mat.Dense) (inverse *mat.Dense) {
    n, _ := m.Dims()
    var q, r, rInv *mat.Dense
    r = mat.NewDense(n, n, nil)
    q = mat.NewDense(n, n, nil)
    rInv = mat.NewDense(n, n, nil)

    inverse = mat.NewDense(n, n, nil)

    qr := mat.QR{}
    qr.Factorize(m)
    qr.QTo(q)
    qr.RTo(r)
    rInv.Inverse(r)

    inverse.Product(rInv, q.T())
    return
}

func nextGuess(c, z, prePredecessor, predecessor *mat.Dense) (guess *mat.Dense) {
    n, _ := c.Dims()
    var doubled, p *mat.Dense
    guess = mat.NewDense(n, n, nil)
    doubled = mat.NewDense(n, n, nil)
    p = mat.NewDense(n, n, nil)
    doubled.Scale(2, predecessor)
    p.Product(z, prePredecessor)
    guess.Add(doubled, p)
    return
}

func isIllConditioned(m* mat.Dense, iteration int) bool {
    n, _ := m.Dims()
    negative := mat.NewDense(n, n, nil)
    negative.Scale(-1, m)
    max := math.Max(mat.Max(m), mat.Max(negative))
    det := mat.Det(m)

    return math.Pow(max, float64(n)) / det > 1e15
}

func Calculate(c *mat.Dense) (sq *mat.Dense, err error) {
    err = nil
    n, _ := c.Dims()
    var m1, m2, m3, eyeN, inverse, p, z *mat.Dense
    eyeN = eye.OfDimension(n)
    sq = mat.NewDense(n, n, nil)
    m1 = mat.NewDense(n, n, nil)
    m2 = mat.NewDense(n, n, nil)
    m3 = mat.NewDense(n, n, nil)
    p = mat.NewDense(n, n, nil)
    m1.Clone(eyeN)
    m2.Clone(c)
    z = mat.NewDense(n, n, nil)
    z.Sub(c, eyeN)

    for i := 1; i <= 100; i++ {
        m3 = nextGuess(c, z, m1, m2)
        m1.Clone(m2)
        m2.Clone(m3)

        if (isIllConditioned(m3, i)) {
            break;
        }
    }

    inverse = inverseViaQR(m1)
    p.Product(m2, inverse)
    sq.Sub(p, eyeN)

    return
}
