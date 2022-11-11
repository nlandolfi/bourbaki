# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make pAC_learnable_hypotheses.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

pAC_learnable_hypotheses.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > pAC_learnable_hypotheses.tex

pAC_learnable_hypotheses.pdf: ../../*.tex ../../trademark.pdf *.tex pAC_learnable_hypotheses.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode pAC_learnable_hypotheses.tex
	../../latexrun pAC_learnable_hypotheses.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: pAC_learnable_hypotheses.pdf
	open pAC_learnable_hypotheses.pdf

o:
	make open

terms: pAC_learnable_hypotheses.tex
	bbk_sheet -mode ts pAC_learnable_hypotheses.tex

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
	rm -f pAC_learnable_hypotheses.*
	make reset
	bbk_sheet -mode mk > Makefile