# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make unordered_triples.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

unordered_triples.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > unordered_triples.tex

unordered_triples.pdf: ../../*.tex ../../trademark.pdf *.tex unordered_triples.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode unordered_triples.tex
	../../latexrun unordered_triples.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: unordered_triples.pdf
	open unordered_triples.pdf

o:
	make open

terms: unordered_triples.tex
	bbk sheet -mode ts unordered_triples.tex

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
	rm -f unordered_triples.*
	make reset
	bbk sheet -mode mk > Makefile