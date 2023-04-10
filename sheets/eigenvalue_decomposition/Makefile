# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make eigenvalue_decomposition.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

eigenvalue_decomposition.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > eigenvalue_decomposition.tex

eigenvalue_decomposition.pdf: ../../*.tex ../../trademark.pdf *.tex eigenvalue_decomposition.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode eigenvalue_decomposition.tex
	../../latexrun eigenvalue_decomposition.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: eigenvalue_decomposition.pdf
	open eigenvalue_decomposition.pdf

o:
	make open

terms: eigenvalue_decomposition.tex
	bbk_sheet -mode ts eigenvalue_decomposition.tex

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
	rm -f eigenvalue_decomposition.*
	make reset
	bbk_sheet -mode mk > Makefile