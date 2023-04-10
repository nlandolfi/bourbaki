# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make differential_equations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

differential_equations.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > differential_equations.tex

differential_equations.pdf: ../../*.tex ../../trademark.pdf *.tex differential_equations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode differential_equations.tex
	../../latexrun differential_equations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: differential_equations.pdf
	open differential_equations.pdf

o:
	make open

terms: differential_equations.tex
	bbk_sheet -mode ts differential_equations.tex

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
	rm -f differential_equations.*
	make reset
	bbk_sheet -mode mk > Makefile