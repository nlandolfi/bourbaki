# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make first_degree_equations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

first_degree_equations.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > first_degree_equations.tex

first_degree_equations.pdf: ../../*.tex ../../trademark.pdf *.tex first_degree_equations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode first_degree_equations.tex
	../../latexrun first_degree_equations.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: first_degree_equations.pdf
	open first_degree_equations.pdf

o:
	make open

terms: first_degree_equations.tex
	bbk sheet -mode ts first_degree_equations.tex

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
	rm -f first_degree_equations.*
	make reset
	bbk sheet -mode mk > Makefile