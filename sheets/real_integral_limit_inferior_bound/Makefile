# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_integral_limit_inferior_bound.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

real_integral_limit_inferior_bound.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_integral_limit_inferior_bound.tex

real_integral_limit_inferior_bound.pdf: ../../*.tex ../../trademark.pdf *.tex real_integral_limit_inferior_bound.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_integral_limit_inferior_bound.tex
	../../latexrun real_integral_limit_inferior_bound.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_integral_limit_inferior_bound.pdf
	open real_integral_limit_inferior_bound.pdf

o:
	make open

terms: real_integral_limit_inferior_bound.tex
	bbk_sheet -mode ts real_integral_limit_inferior_bound.tex

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
	rm -f real_integral_limit_inferior_bound.*
	make reset
	bbk_sheet -mode mk > Makefile