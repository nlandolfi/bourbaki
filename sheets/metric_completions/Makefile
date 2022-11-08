# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make metric_completions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

metric_completions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > metric_completions.tex

metric_completions.pdf: ../../*.tex ../../trademark.pdf *.tex metric_completions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode metric_completions.tex
	../../latexrun metric_completions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: metric_completions.pdf
	open metric_completions.pdf

o:
	make open

terms: metric_completions.tex
	bbk_sheet -mode ts metric_completions.tex

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
	rm -f metric_completions.*
	make reset
	bbk_sheet -mode mk > Makefile