# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make eigenvalues_and_eigenvectors.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

eigenvalues_and_eigenvectors.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > eigenvalues_and_eigenvectors.tex

eigenvalues_and_eigenvectors.pdf: ../../*.tex ../../trademark.pdf *.tex eigenvalues_and_eigenvectors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode eigenvalues_and_eigenvectors.tex
	../../latexrun eigenvalues_and_eigenvectors.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: eigenvalues_and_eigenvectors.pdf
	open eigenvalues_and_eigenvectors.pdf

o:
	make open

terms: eigenvalues_and_eigenvectors.tex
	bbk sheet -mode ts eigenvalues_and_eigenvectors.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f eigenvalues_and_eigenvectors.*
	make reset
	bbk sheet -mode mk > Makefile