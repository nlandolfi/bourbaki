# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make sigma_algebras.pdf

sigma_algebras.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > sigma_algebras.tex

sigma_algebras.pdf: ../../*.tex ../../trademark.pdf *.tex sigma_algebras.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode sigma_algebras.tex
	../../latexrun sigma_algebras.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: sigma_algebras.pdf
	open sigma_algebras.pdf

o:
	make open

terms: sigma_algebras.tex
	bbk_sheet -mode ts sigma_algebras.tex

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
	rm -f sigma_algebras.*
	make reset
	bbk_sheet -mode mk > Makefile