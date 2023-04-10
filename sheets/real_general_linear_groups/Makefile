# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_general_linear_groups.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

real_general_linear_groups.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > real_general_linear_groups.tex

real_general_linear_groups.pdf: ../../*.tex ../../trademark.pdf *.tex real_general_linear_groups.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_general_linear_groups.tex
	../../latexrun real_general_linear_groups.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_general_linear_groups.pdf
	open real_general_linear_groups.pdf

o:
	make open

terms: real_general_linear_groups.tex
	bbk_sheet -mode ts real_general_linear_groups.tex

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
	rm -f real_general_linear_groups.*
	make reset
	bbk_sheet -mode mk > Makefile