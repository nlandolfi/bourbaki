# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make adjoints_of_linear_transformations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

adjoints_of_linear_transformations.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > adjoints_of_linear_transformations.tex

adjoints_of_linear_transformations.pdf: ../../*.tex ../../trademark.pdf *.tex adjoints_of_linear_transformations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode adjoints_of_linear_transformations.tex
	../../latexrun adjoints_of_linear_transformations.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: adjoints_of_linear_transformations.pdf
	open adjoints_of_linear_transformations.pdf

o:
	make open

terms: adjoints_of_linear_transformations.tex
	bbk sheet -mode ts adjoints_of_linear_transformations.tex

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
	rm -f adjoints_of_linear_transformations.*
	make reset
	bbk sheet -mode mk > Makefile