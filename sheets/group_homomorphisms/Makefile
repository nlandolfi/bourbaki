# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make group_homomorphisms.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

group_homomorphisms.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > group_homomorphisms.tex

group_homomorphisms.pdf: ../../*.tex ../../trademark.pdf *.tex group_homomorphisms.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode group_homomorphisms.tex
	../../latexrun group_homomorphisms.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: group_homomorphisms.pdf
	open group_homomorphisms.pdf

o:
	make open

terms: group_homomorphisms.tex
	bbk sheet -mode ts group_homomorphisms.tex

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
	rm -f group_homomorphisms.*
	make reset
	bbk sheet -mode mk > Makefile