# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make metric_convergence.pdf

metric_convergence.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > metric_convergence.tex

metric_convergence.pdf: ../../*.tex ../../trademark.pdf *.tex metric_convergence.tex
	# pdflatex --file-line-error -interaction=nonstopmode metric_convergence.tex
	../../latexrun metric_convergence.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: metric_convergence.pdf
	open metric_convergence.pdf

o:
	make open

terms: metric_convergence.tex
	bbk_sheet -mode ts metric_convergence.tex

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
	rm -f metric_convergence.*
	make reset
	bbk_sheet -mode mk > Makefile