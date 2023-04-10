# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make directed_shortest_path_problems.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

directed_shortest_path_problems.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > directed_shortest_path_problems.tex

directed_shortest_path_problems.pdf: ../../*.tex ../../trademark.pdf *.tex directed_shortest_path_problems.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode directed_shortest_path_problems.tex
	../../latexrun directed_shortest_path_problems.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: directed_shortest_path_problems.pdf
	open directed_shortest_path_problems.pdf

o:
	make open

terms: directed_shortest_path_problems.tex
	bbk_sheet -mode ts directed_shortest_path_problems.tex

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
	rm -f directed_shortest_path_problems.*
	make reset
	bbk_sheet -mode mk > Makefile